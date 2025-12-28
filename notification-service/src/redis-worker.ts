import Redis from "ioredis";
import { config } from "./config";
import { EmailJob } from "./types";
import { sendVerificationEmail, sendWelcomeEmail } from "./email";

const redis = new Redis(config.redisUrl, {
  maxRetriesPerRequest: null,
});

let isProcessing = false;

async function processJob(jobId: string): Promise<void> {
  try {
    // Get job data from hash
    const jobData = await redis.hgetall(`${config.queueName}:${jobId}`);

    if (!jobData || !jobData.data) {
      console.error(`Job ${jobId} not found or has no data`);
      return;
    }

    const job: EmailJob = JSON.parse(jobData.data);
    console.log(`Processing ${job.type} email job (ID: ${jobId})`);

    switch (job.type) {
      case "verification":
        await sendVerificationEmail(job);
        break;
      case "welcome":
        await sendWelcomeEmail(job);
        break;
      default:
        const exhaustiveCheck: never = job;
        console.warn(`Unknown email type: ${(exhaustiveCheck as any).type}`);
    }

    // Remove job from hash after successful processing
    await redis.del(`${config.queueName}:${jobId}`);
    console.log(`✓ Job ${jobId} completed successfully`);
  } catch (error) {
    console.error(`✗ Job ${jobId} failed:`, error);
    throw error;
  }
}

async function pollQueue(): Promise<void> {
  if (isProcessing) return;

  try {
    isProcessing = true;

    // Pop job from wait list (blocking pop with 5 second timeout)
    const result = await redis.brpop(`${config.queueName}:wait`, 5);

    if (result) {
      const [, jobId] = result;
      await processJob(jobId);
    }
  } catch (error) {
    console.error("Queue polling error:", error);
  } finally {
    isProcessing = false;
    // Continue polling
    setImmediate(() => pollQueue());
  }
}

export function startRedisWorker(): void {
  console.log(
    `📧 Redis worker started, listening on queue: ${config.queueName}`
  );

  // Start polling
  pollQueue();

  // Handle graceful shutdown
  process.on("SIGTERM", async () => {
    console.log("Shutting down Redis worker...");
    await redis.quit();
    process.exit(0);
  });
}
