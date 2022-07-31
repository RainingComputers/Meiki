<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { formatRequestError } from "$lib/api/request"
    import { logout } from "$lib/api/user"
    import Button, { ButtonType } from "$cmp/Button.svelte"
    import Error from "$cmp/toast/Error.svelte"
    import Logo from "$cmp/app/Logo.svelte"
    import Window from "$cmp/modal/Window.svelte"

    const dispatchEvent = createEventDispatcher()

    let error = ""

    async function logoutUser() {
        try {
            await logout()
            dispatchEvent("loggedOut")
        } catch (err) {
            error = formatRequestError(err, "logging out")
        }
    }
</script>

<Window width="500px">
    <div />
    <Logo />
    <div class="px-4 py-5 flex flex-col gap-5">
        {#if error}
            <Error>{error}</Error>
        {/if}
        <Button label="Logout" type={ButtonType.SECONDARY} onClick={logoutUser} />
    </div>

    <div class="flex flex-col text-center text-sm mt-3 text-content">
        <div>The Meiki notes editor</div>
        <div>&#169; Copyright 2022 Vishnu Shankar B and Alex Joseph</div>
    </div>
</Window>
