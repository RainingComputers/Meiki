<script lang="ts">
    import Button from "$cmp/Button.svelte"
    import Card from "$cmp/Card.svelte"
    import Entry from "$cmp/Entry.svelte"
    import Logo from "$cmp/app/Logo.svelte"
    import Error from "$cmp/toast/Error.svelte"

    export let error: string
    export let buttonLabel: string
    export let onClick: () => void
    export let confirmPassword: boolean = false

    let usernameEntry: Entry
    let passwordEntry: Entry
    let confirmPasswordEntry: Entry

    export function getUsername() {
        return usernameEntry.getValue()
    }

    export function getPassword() {
        return passwordEntry.getValue()
    }

    export function getConfirmPassword() {
        return confirmPasswordEntry.getValue()
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
            <Entry label="Username" bind:this={usernameEntry} onEnter={onClick} />
            <Entry label="Password" bind:this={passwordEntry} password={true} onEnter={onClick} />
            {#if confirmPassword}
                <Entry
                    label="Confirm password"
                    bind:this={confirmPasswordEntry}
                    password={true}
                    onEnter={onClick}
                />
            {/if}
        </div>

        <Button {onClick} fullWidth={true} label={buttonLabel} />
    </div>
</Card>
