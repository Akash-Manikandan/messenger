import { userClient } from '$lib/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params }) => {
	const userId = '01KDQFAAYVVV51Y2626CH3Q22V';

	const user = await userClient.getUser({ id: userId });
	return { user };
};
