/** @jsxImportSource preact */
import render from "preact-render-to-string";
import { sendEmail } from "./mailer";
import { VerificationEmailJob, WelcomeEmailJob } from "./types";
import {
  VerificationEmail,
  getVerificationEmailText,
} from "./templates/VerificationEmail";
import { WelcomeEmail, getWelcomeEmailText } from "./templates/WelcomeEmail";

export async function sendVerificationEmail(
  job: VerificationEmailJob
): Promise<void> {
  const { to, fromEmail, fromName, data } = job;

  console.log(`Sending verification email to ${to}`);

  const html = render(
    <VerificationEmail username={data.username} url={data.url} />
  );
  const text = getVerificationEmailText(data.username, data.url);

  try {
    await sendEmail({
      to,
      from: `${fromName} <${fromEmail}>`,
      subject: "Verify your email address",
      text,
      html: `<!DOCTYPE html>${html}`,
    });

    console.log(`✓ Verification email sent to ${to}`);
  } catch (error) {
    console.error(`✗ Failed to send verification email to ${to}:`, error);
    throw error;
  }
}

export async function sendWelcomeEmail(job: WelcomeEmailJob): Promise<void> {
  const { to, fromEmail, fromName, data } = job;

  console.log(`Sending welcome email to ${to}`);

  const html = render(<WelcomeEmail username={data.username} />);
  const text = getWelcomeEmailText(data.username);

  try {
    await sendEmail({
      to,
      from: `${fromName} <${fromEmail}>`,
      subject: "Welcome to Messenger!",
      text,
      html: `<!DOCTYPE html>${html}`,
    });

    console.log(`✓ Welcome email sent to ${to}`);
  } catch (error) {
    console.error(`✗ Failed to send welcome email to ${to}:`, error);
    throw error;
  }
}
