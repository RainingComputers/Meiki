<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { formatRequestError } from "$lib/api/request"
    import { logout } from "$lib/api/user"
    import Button, { ButtonType } from "$cmp/Button.svelte"
    import Error from "$cmp/toast/Error.svelte"
    import Logo from "../Logo.svelte"

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

<div
    class=" w-[500px] max-h-fit m-auto bg-gray-50 border-gray-200 rounded-xl flex flex-col py-5 px-6"
>
    <div />
    <Logo />
    <div class="px-4 py-5 flex flex-col gap-5">
        {#if error}
            <Error>{error}</Error>
        {/if}
        <Button
            label="Logout"
            type={ButtonType.SECONDARY}
            onClick={logoutUser}
        />
    </div>

    <div class="flex flex-col text-center text-sm mt-3">
        <div class="">The Meiki notes editor</div>
        <div class="">
            &#169; Copyright 2022 Vishnu Shankar B and Alex Joseph
        </div>
    </div>
</div>
