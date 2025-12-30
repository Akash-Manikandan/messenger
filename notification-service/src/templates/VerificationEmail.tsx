/** @jsxImportSource preact */
import { EmailLayout } from "./EmailLayout";
import { EmailSignature } from "./EmailSignature";

interface VerificationEmailProps {
  username: string;
  url: string;
}

export function VerificationEmail({ username, url }: VerificationEmailProps) {
  return (
    <EmailLayout>
      <div>
        <h1 style="font-size: 24px; font-weight: bold; margin-top: 0; color: #333;">
          Verify your email address
        </h1>
        <p style="font-size: 16px; line-height: 1.6; color: #555;">
          Hello <strong>{username}</strong>,
        </p>
        <p style="font-size: 16px; line-height: 1.6; color: #555;">
          Thank you for signing up! Please verify your email address to get
          started.
        </p>
        <div style="text-align: center; margin: 32px 0;">
          <a
            href={url}
            style="display: inline-block; background-color: #007bff; color: white; padding: 12px 32px; text-decoration: none; border-radius: 6px; font-weight: bold; font-size: 16px;"
          >
            Verify Email
          </a>
        </div>
        <div style="background-color: #f8f9fa; padding: 20px; border-radius: 6px; margin: 24px 0;">
          <p style="margin: 0; font-size: 14px; color: #666;">
            Or copy and paste this link in your browser:
          </p>
          <p style="margin: 8px 0 0 0; font-size: 14px; color: #007bff; word-break: break-all;">
            {url}
          </p>
        </div>
        <p style="font-size: 14px; color: #888; margin-top: 32px;">
          If you didn't sign up for this account, you can safely ignore this
          email.
        </p>
        <EmailSignature />
      </div>
    </EmailLayout>
  );
}

export function getVerificationEmailText(
  username: string,
  url: string
): string {
  return `Hello ${username},

Thank you for signing up! Please verify your email address to get started.

Verification link: ${url}

If you didn't sign up for this account, you can safely ignore this email.`;
}
