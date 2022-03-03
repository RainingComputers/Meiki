<script lang="ts">
    import Root from "../components/Root.svelte"
    import CredentialForm, {
        CredentialFromType,
    } from "../appComponents/CredentialForm.svelte"
</script>

<Root>
    <CredentialForm type={CredentialFromType.CREATE} />
</Root>
