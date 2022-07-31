<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"
    import { tokensPresent } from "$lib/api/user"
    import Root from "$cmp/Root.svelte"
    import Logo from "$cmp/app/Logo.svelte"
    import Button, { ButtonType } from "$cmp/Button.svelte"
    import Link from "$cmp/Link.svelte"

    let loggedIn: boolean = true

    onMount(() => {
        loggedIn = tokensPresent()
        if (loggedIn) goto("/notes")
    })
</script>

{#if !loggedIn}
    <Root>
        <span class="flex flex-col items-center gap-6">
            <Logo width="50%" />
            <h1 class="text-5xl">The ⚡ lightning fast ⚡ notes editor</h1>
            <h1 class="text-2xl">
                Meiki is a markdown notes editor built with performance and simplicity in mind
            </h1>
        </span>
        <span class="flex flex-col gap-4 p-14">
            <Button
                type={ButtonType.HERO}
                label="Create Meiki Account"
                onClick={() => {
                    goto("/create")
                }}
            />
            <span class="text-lg">
                Already have an account? <Link text="Click here" target="/login" /> to login
            </span>
        </span>
    </Root>
{/if}
