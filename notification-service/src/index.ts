import { startRedisWorker } from "./redis-worker";

console.log("🚀 Starting Notification Service...");

startRedisWorker();

// Graceful shutdown
process.on("SIGTERM", () => {
  console.log("\n👋 Shutting down gracefully...");
  process.exit(0);
});

process.on("SIGINT", () => {
  console.log("\n👋 Shutting down gracefully...");
  process.exit(0);
});
