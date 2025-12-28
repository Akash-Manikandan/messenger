/** @jsxImportSource preact */
import { EmailLayout } from "./EmailLayout";

interface WelcomeEmailProps {
  username: string;
}

export function WelcomeEmail({ username }: WelcomeEmailProps) {
  return (
    <EmailLayout>
      <div>
        <h1 style="font-size: 28px; font-weight: bold; margin-top: 0; color: #333;">
          Welcome to Messenger! 🎉
        </h1>
        <p style="font-size: 16px; line-height: 1.6; color: #555;">
          Hello <strong>{username}</strong>,
        </p>
        <p style="font-size: 16px; line-height: 1.6; color: #555;">
          We're thrilled to have you join our community! Messenger makes it easy
          to stay connected with friends and family.
        </p>
        <div style="margin: 32px 0;">
          <img
            src="https://assets-examples.mailtrap.io/integration-examples/welcome.png"
            alt="Welcome"
            style="width: 100%; border-radius: 8px;"
          />
        </div>
        <p style="font-size: 16px; line-height: 1.6; color: #555;">
          Start exploring and connecting today!
        </p>
        <p style="font-size: 14px; color: #888; margin-top: 32px; border-top: 1px solid #eee; padding-top: 20px;">
          Best regards,
          <br />
          The Messenger Team
        </p>
      </div>
    </EmailLayout>
  );
}

export function getWelcomeEmailText(username: string): string {
  return `Welcome to Messenger! 🎉

Hello ${username},

We're thrilled to have you join our community! Messenger makes it easy to stay connected with friends and family.

Start exploring and connecting today!

Best regards,
The Messenger Team`;
}
