<script lang="ts">
    import Root from "../components/Root.svelte"
    import CreateUserForm from "../appComponents/userForm/CreateUserForm.svelte"
    import UserCreated from "../appComponents/userForm/UserCreated.svelte"

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
