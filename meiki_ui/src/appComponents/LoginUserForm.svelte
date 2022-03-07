<script lang="ts">
    // TODO: Add input validation
    // TODO: show user exists

    import { createEventDispatcher } from "svelte"
    import { login } from "../api/login"
    import Card from "../components/Card.svelte"
    import Button from "../components/Button.svelte"
    import Entry from "../components/Entry.svelte"
    import Error from "../components/Error.svelte"
    import Logo from "./Logo.svelte"

    let usernameEntry: Entry
    let passwordEntry: Entry
    let error: boolean = false
    const dispatch = createEventDispatcher()

    async function onClick() {
        error = false

        try {
            await login(usernameEntry.getValue(), passwordEntry.getValue())

            dispatch("userLoggedIn")
        } catch (err) {
            console.log(err)
            error = true
        }
    }
</script>

<Card>
    <div class="flex flex-col gap-10 items-center p-5">
        <Logo />
        <div class="flex flex-col w-full gap-4">
            {#if error}
                <Error>An error has occured while logging in, try again</Error>
            {/if}
            <Entry label="Username" bind:this={usernameEntry} />
            <Entry label="Password" bind:this={passwordEntry} password={true} />
        </div>

        <Button {onClick} expand={true} label="Login" />
    </div>
</Card>
