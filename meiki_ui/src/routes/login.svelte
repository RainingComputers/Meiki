<script lang="ts">
    import Root from "../components/Root.svelte"
    import Link from "../components/Link.svelte"
    import CredentialForm from "../appComponents/CredentialForm.svelte"
</script>

<Root>
    <CredentialForm />
    <div class=" mt-8 text-sm">
        New to Meiki notes? <Link text="Create" target="/create" /> a new account
        to get started
    </div>
</Root>
