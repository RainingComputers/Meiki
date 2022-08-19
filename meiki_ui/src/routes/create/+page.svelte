<script lang="ts">
    import Root from "$cmp/Root.svelte"
    import CreateUserForm from "$cmp/app/userForm/CreateUserForm.svelte"
    import UserCreated from "$cmp/app/userForm/UserCreated.svelte"

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
