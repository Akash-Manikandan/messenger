import { BaseEmailJob } from "./base";

// Welcome email job data
export interface WelcomeEmailData {
  username: string;
}

// Welcome email job
export interface WelcomeEmailJob extends BaseEmailJob {
  type: "welcome";
  data: WelcomeEmailData;
}
