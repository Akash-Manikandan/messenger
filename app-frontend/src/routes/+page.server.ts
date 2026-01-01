import { userClient } from '$lib/api';
import type { PageServerLoad } from './$types';
import { createAvatar } from '@dicebear/core';
import { funEmoji } from '@dicebear/collection';

export const load: PageServerLoad = async () => {
	const data = await userClient.listUsers({ limit: 20, offset: 0 });
	for (const user of data.users) {
		user.avatar = createAvatar(funEmoji, {
			seed: user.avatar,
			size: 36,
			scale: 90,
			backgroundType: ['gradientLinear', 'solid'],
			eyes: ['cute', 'glasses', 'love', 'shades', 'wink', 'wink2', 'plain'],
			mouth: ['lilSmile', 'wideSmile', 'smileTeeth', 'smileLol', 'kissHeart']
		}).toDataUri();
	}

	return { users: data.users };
};
