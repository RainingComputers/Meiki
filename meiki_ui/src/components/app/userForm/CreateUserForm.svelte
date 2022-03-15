<script lang="ts">
    // TODO: Add input validation
    // TODO: show user exists
    // TODO: Better error messages

    import { createEventDispatcher } from "svelte"
    import { createAccount } from "$lib/api/createAccount"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        try {
            await createAccount(userForm.getUsername(), userForm.getPassword())
            dispatch("userCreated")
        } catch (err) {
            error =
                "An error has occured while creating the account, please try again later"
        }
    }
</script>

<UserForm
    bind:this={userForm}
    {error}
    {onClick}
    buttonLabel={"Create Meiki account"}
/>
