import { createClient } from '@connectrpc/connect';
import { createConnectTransport } from '@connectrpc/connect-web';
import { UserService } from '$lib/proto/user_pb';
import { HealthService } from '$lib/proto/health_pb';

// Get the backend URL from environment variables or use default
const BACKEND_URL = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080';

// Create a Connect transport for browser
const transport = createConnectTransport({
	baseUrl: BACKEND_URL,
	// Add any interceptors or custom options here
	useBinaryFormat: true // Use binary format for better performance
});

// Create typed clients using service descriptors from protoc-gen-es v2
export const userClient = createClient(UserService, transport);

export const healthClient = createClient(HealthService, transport);

// Export transport for custom usage if needed
export { transport };
