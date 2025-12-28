// Base interface for all email jobs
export interface BaseEmailJob {
  to: string;
  fromEmail: string;
  fromName: string;
}
