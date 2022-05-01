<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { createUser } from "$lib/api/user"
    import { StatusNotOkError } from "$lib/api/request"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        if (userForm.getPassword() != userForm.getConfirmPassword()) {
            error = "Passwords do not match"
            return
        }

        try {
            await createUser(userForm.getUsername(), userForm.getPassword())
            dispatch("userCreated")
        } catch (err) {
            if (err instanceof StatusNotOkError) {
                error = await err.message
                return
            }

            error =
                "An error has occured while creating the account, unable to connet to server"
        }
    }
</script>

<UserForm
    bind:this={userForm}
    {error}
    {onClick}
    buttonLabel={"Create Meiki account"}
    confirmPassword={true}
/>
