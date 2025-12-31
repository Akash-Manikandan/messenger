<script lang="ts">
    let { data } = $props();

    let user = $derived(data.user);
</script>

<h1>User Profile</h1>
{#if user}
    <p><strong>ID:</strong> {user.id}</p>
    <p><strong>Name:</strong> {user.username}</p>
    <p><strong>Email:</strong> {user.email}</p>
{:else}
    <p>Loading user data...</p>
{/if}