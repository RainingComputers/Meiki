<script lang="ts">
    import { goto } from "$app/navigation"
    import Button, { ButtonType } from "$cmp/Button.svelte"
    import Logo from "./Logo.svelte"
    import { logout } from "$lib/api/user"

    export async function handleLogout() {
        await logout()
        goto("/login")
    }
</script>

<div
    class=" w-[600px] h-96 m-auto bg-gray-50 border-gray-200 rounded-xl flex flex-col py-0 px-6 gap-8"
>
    <div />
    <Logo />
    <div class="px-5 flex flex-col">
        <Button
            label="Logout"
            type={ButtonType.SECONDARY}
            onClick={handleLogout}
        />
    </div>

    <div class="flex flex-col text-center text-sm">
        <div class="">The Meiki notes editor</div>
        <div class="">
            &#169; Copyright 2022 Vishnu Shankar B and Alex Joseph
        </div>
    </div>
</div>
