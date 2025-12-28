import nodemailer from "nodemailer";
import { config } from "./config";

export const transporter = nodemailer.createTransport({
  host: config.smtp.host,
  port: config.smtp.port,
  secure: false, // Use TLS
  auth: {
    user: config.smtp.user,
    pass: config.smtp.password,
  },
});

export interface SendEmailOptions {
  to: string;
  from: string;
  subject: string;
  text: string;
  html: string;
}

export async function sendEmail(options: SendEmailOptions): Promise<void> {
  await transporter.sendMail({
    from: options.from,
    to: options.to,
    subject: options.subject,
    text: options.text,
    html: options.html,
  });
}
