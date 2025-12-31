// Export all API services
export { userService } from './user-service';
export { healthService } from './health-service';
export { userClient, healthClient, transport } from './grpc-client';

// Re-export types for convenience
export type * from '$lib/proto/user_pb';
export type * from '$lib/proto/health_pb';
