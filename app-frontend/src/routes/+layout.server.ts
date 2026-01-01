import { userClient } from '$lib/api';
import type { PageServerLoad } from './$types';
import { createAvatar } from '@dicebear/core';
import { funEmoji } from '@dicebear/collection';

export const load: PageServerLoad = async () => {
	const userId = '01KDQFAAYVVV51Y2626CH3Q22V';
	const user = await userClient.getUser({ id: userId });
	user.avatar = createAvatar(funEmoji, {
		seed: user.avatar,
		size: 36,
		scale: 90,
		backgroundType: ['gradientLinear', 'solid'],
		eyes: ['cute', 'glasses', 'love', 'shades', 'wink', 'wink2', 'plain'],
		mouth: ['lilSmile', 'wideSmile', 'smileTeeth', 'smileLol', 'kissHeart']
	}).toDataUri();
	return { user };
};
