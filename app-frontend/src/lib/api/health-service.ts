import { healthClient } from './grpc-client';

/**
 * Health service wrapper for checking backend service health
 */
export const healthService = {
	/**
	 * Check the health status of the backend service
	 */
	async check() {
		return await healthClient.check({});
	}
};
