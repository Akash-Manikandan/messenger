<script lang="ts">
	import { ScrollArea } from '$lib/components/ui/scroll-area';
	import * as Tooltip from '$lib/components/ui/tooltip';
	let { data } = $props();
	let users = $derived(data.users);
</script>

<section class="m-6 flex flex-col gap-4">
	<h1>People you may know</h1>
	<ScrollArea class="w-full rounded-md border" orientation="horizontal">
		<div class="flex w-max gap-4 px-3 pt-2 pb-3">
			{#each users as user}
				<Tooltip.Root>
					<Tooltip.Trigger>
						<figure class="flex shrink-0 flex-col items-center gap-2">
							<img
								src={user.avatar}
								alt={user.username}
								width={48}
								height={48}
								class="h-12 w-12 rounded-full"
							/>
							<figcaption class="text-sm">{user.username}</figcaption>
						</figure>
					</Tooltip.Trigger>
					<Tooltip.Content>
						<p>{user.username}</p>
					</Tooltip.Content>
				</Tooltip.Root>
			{/each}
		</div>
	</ScrollArea>
</section>
