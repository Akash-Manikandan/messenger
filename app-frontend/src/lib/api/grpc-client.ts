import { createClient } from '@connectrpc/connect';
import { createGrpcTransport } from '@connectrpc/connect-node';
import { UserService } from '$lib/proto/user_pb';
import { HealthService } from '$lib/proto/health_pb';

// Get the backend URL from environment variables or use default
const BACKEND_URL = import.meta.env.VITE_GRPC_URL || 'http://localhost:9090';

// Create a gRPC transport for Node.js server to connect to your Go gRPC backend
const transport = createGrpcTransport({
	baseUrl: BACKEND_URL,
	idleConnectionTimeoutMs: 30000,
	useBinaryFormat: true
});

// Create typed clients using service descriptors from protoc-gen-es v2
export const userClient = createClient(UserService, transport);

export const healthClient = createClient(HealthService, transport);

// Export transport for custom usage if needed
export { transport };
