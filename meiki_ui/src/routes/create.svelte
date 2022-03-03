<script lang="ts">
    import Root from "../components/Root.svelte"
    import CredentialForm, {
        CredentialFromType,
    } from "../appComponents/CredentialForm.svelte"
    import UserCreated from "../appComponents/UserCreated.svelte"

    let userCreated: boolean = false
</script>

<Root>
    {#if userCreated}
        <UserCreated />
    {:else}
        <CredentialForm
            type={CredentialFromType.CREATE}
            on:userCreated={() => {
                userCreated = true
            }}
        />
    {/if}
</Root>
