/** @jsxImportSource preact */
import { VNode } from "preact";

interface EmailLayoutProps {
  children: VNode;
}

export function EmailLayout({ children }: EmailLayoutProps) {
  return (
    <html>
      <head>
        <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
      </head>
      <body style="font-family: sans-serif; margin: 0; padding: 0; background-color: #f4f4f4;">
        <div style="display: block; margin: 20px auto; max-width: 600px; background-color: white; padding: 40px; border-radius: 8px;">
          {children}
        </div>
      </body>
    </html>
  );
}
