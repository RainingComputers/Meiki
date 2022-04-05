<script lang="ts">
    // TODO: Add input validation

    import { createEventDispatcher } from "svelte"
    import { createAccount } from "$lib/api/createAccount"
    import { StatusNotOkError } from "$lib/api/request"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        try {
            await createAccount(userForm.getUsername(), userForm.getPassword())
            dispatch("userCreated")
        } catch (err) {
            if (err instanceof StatusNotOkError) {
                error = await err.message
                return
            }

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
