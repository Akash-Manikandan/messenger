import { BaseEmailJob } from "./base";

// Verification email job data
export interface VerificationEmailData {
  username: string;
  url: string;
}

// Verification email job
export interface VerificationEmailJob extends BaseEmailJob {
  type: "verification";
  data: VerificationEmailData;
}
