import { userClient } from './grpc-client';

/**
 * User service wrapper providing convenient methods for user operations
 *
 * Note: Using @ts-ignore due to type incompatibility between
 * @connectrpc/connect v1.6.1 and @bufbuild/protobuf v1.10.0
 */
export const userService = {
	/**
	 * Create a new user
	 */
	async createUser(data: {
		username: string;
		email: string;
		password: string;
		firstName?: string;
		lastName?: string;
	}) {
		// @ts-ignore
		return await userClient.createUser({
			username: data.username,
			email: data.email,
			password: data.password,
			firstName: data.firstName,
			lastName: data.lastName
		});
	},

	/**
	 * Get a user by ID
	 */
	async getUser(id: string) {
		// @ts-ignore
		return await userClient.getUser({ id });
	},

	/**
	 * Update an existing user
	 */
	async updateUser(data: {
		id: string;
		firstName?: string;
		lastName?: string;
		avatar?: string;
		bio?: string;
		isActive?: boolean;
		isVerified?: boolean;
	}) {
		// @ts-ignore
		return await userClient.updateUser({
			id: data.id,
			firstName: data.firstName,
			lastName: data.lastName,
			avatar: data.avatar,
			bio: data.bio,
			isActive: data.isActive,
			isVerified: data.isVerified
		});
	},

	/**
	 * Delete a user by ID
	 */
	async deleteUser(id: string) {
		// @ts-ignore
		return await userClient.deleteUser({ id });
	},

	/**
	 * List users with pagination
	 */
	async listUsers(params?: { limit?: number; offset?: number }) {
		// @ts-ignore
		return await userClient.listUsers({
			limit: params?.limit || 10,
			offset: params?.offset || 0
		});
	},

	/**
	 * Verify a user's email with user ID
	 */
	async verifyUser(userId: string) {
		// @ts-ignore
		return await userClient.verifyUser({ userId });
	}
};
