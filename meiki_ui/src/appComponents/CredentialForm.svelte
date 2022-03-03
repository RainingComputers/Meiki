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

    import { createEventDispatcher } from "svelte"
    import { createAccount } from "../api/createAccount"
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
    import Entry from "../components/Entry.svelte"
    import Error from "../components/Error.svelte"
    import Logo from "./Logo.svelte"

    export let type: CredentialFromType = CredentialFromType.LOGIN

    let usernameEntry: Entry
    let passwordEntry: Entry
    let error: boolean
    const dispatch = createEventDispatcher()

    async function onClick() {
        error = false

        try {
            await createAccount(usernameEntry.getValue(), passwordEntry.getValue())
            dispatch("userCreated")
        } catch {
            error = true
        }
    }
</script>

<Card>
    <div class="flex flex-col gap-10 items-center p-5">
        <Logo />
        <div class="flex flex-col w-full gap-4">
            {#if error}
                <Error>
                    An error has occured while creating account, please try
                    again later.
                </Error>
            {/if}
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
