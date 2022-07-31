<script lang="ts">
    import { goto } from "$app/navigation"
    import Root from "$cmp/Root.svelte"
    import Link from "$cmp/Link.svelte"
    import LoginUserForm from "$cmp/app/userForm/LoginUserForm.svelte"
</script>

<Root>
    <LoginUserForm
        on:userLoggedIn={() => {
            goto("/notes")
        }}
    />
    <div class=" mt-8 text-sm text-content">
        New to Meiki notes? <Link text="Create" target="/create" /> a new account to get started
    </div>
</Root>
