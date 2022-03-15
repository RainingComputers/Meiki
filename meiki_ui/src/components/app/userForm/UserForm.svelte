<script lang="ts">
    import Button from "$cmp/Button.svelte"

    import Card from "$cmp/Card.svelte"
    import Entry from "$cmp/Entry.svelte"
    import Logo from "../Logo.svelte"
    import Error from "$cmp/toast/Error.svelte"

    export let error: string
    export let buttonLabel: string
    export let onClick: any

    let usernameEntry: Entry
    let passwordEntry: Entry

    export function getUsername() {
        return usernameEntry.getValue()
    }

    export function getPassword() {
        return passwordEntry.getValue()
    }
</script>

<Card>
    <div class="flex flex-col gap-10 items-center p-5">
        <Logo />
        <div class="flex flex-col w-full gap-4">
            {#if error}
                <Error fullWidth={true}>
                    {error}
                </Error>
            {/if}
            <Entry label="Username" bind:this={usernameEntry} />
            <Entry label="Password" bind:this={passwordEntry} password={true} />
        </div>

        <Button {onClick} fullWidth={true} label={buttonLabel} />
    </div>
</Card>
