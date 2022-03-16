<script lang="ts">
    import Root from "$cmp/Root.svelte"
    import Link from "$cmp/Link.svelte"
    import LoginUserForm from "$cmp/app/userForm/LoginUserForm.svelte"
    import { goto } from "$app/navigation"
</script>

<Root>
    <LoginUserForm
        on:userLoggedIn={() => {
            goto("/notes")
        }}
    />
    <div class=" mt-8 text-sm">
        New to Meiki notes? <Link text="Create" target="/create" /> a new account
        to get started
    </div>
</Root>
