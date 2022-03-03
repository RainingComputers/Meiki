<script lang="ts">
    import Root from "../components/Root.svelte"
    import CreateUserForm from "../appComponents/CreateUserForm.svelte"
    import UserCreated from "../appComponents/UserCreated.svelte"

    let userCreated: boolean = false
</script>

<Root>
    {#if userCreated}
        <UserCreated />
    {:else}
        <CreateUserForm
            on:userCreated={() => {
                userCreated = true
            }}
        />
    {/if}
</Root>
