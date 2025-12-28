import dotenv from "dotenv";

dotenv.config();

export const config = {
  redisUrl: process.env.REDIS_URL!,
  mailtrapToken: process.env.MAIL_TRAP!,
  queueName: "email-jobs",
  smtp: {
    host: process.env.SMTP_HOST || "live.smtp.mailtrap.io",
    port: parseInt(process.env.SMTP_PORT || "587"),
    user: process.env.SMTP_USER || "api",
    password: process.env.MAIL_TRAP!,
    fromEmail: process.env.SMTP_FROM_EMAIL || "hello@demomailtrap.co",
    fromName: process.env.SMTP_FROM_NAME || "Messenger",
  },
};
