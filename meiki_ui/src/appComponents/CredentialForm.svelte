<script context="module" lang="ts">
    export enum CredentialFromType {
        LOGIN,
        CREATE,
    }
</script>

<script lang="ts">
    // TODO: Add input validation
    // TODO: Add toast
    // TODO: redirect to login

    import { createAccount } from "../api/createAccount"
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
    import Entry from "../components/Entry.svelte"
    import Logo from "./Logo.svelte"

    export let type: CredentialFromType = CredentialFromType.LOGIN

    let usernameEntry: Entry
    let passwordEntry: Entry

    function onClick() {
        // TODO: Check for errors
        createAccount(usernameEntry.getValue(), passwordEntry.getValue())
    }
</script>

<Card>
    <div class="flex flex-col gap-10 items-center p-5">
        <Logo />
        <div class="flex flex-col w-full gap-4">
            <Entry label="Username" bind:this={usernameEntry} />
            <Entry label="Password" bind:this={passwordEntry} password={true} />
        </div>
        {#if type === CredentialFromType.LOGIN}
            <Button expand={true} label="Login" />
        {:else if type === CredentialFromType.CREATE}
            <Button {onClick} expand={true} label="Create Meiki account" />
        {/if}
    </div>
</Card>
