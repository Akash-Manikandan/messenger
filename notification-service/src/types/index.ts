// Export all types
export * from "./base";
export * from "./verification";
export * from "./welcome";

// Import for discriminated union
import { VerificationEmailJob } from "./verification";
import { WelcomeEmailJob } from "./welcome";

// Discriminated union of all email job types
export type EmailJob = VerificationEmailJob | WelcomeEmailJob;

// Extract the type for convenience
export type EmailJobType = EmailJob["type"];
