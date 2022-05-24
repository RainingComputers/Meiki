<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { login } from "$lib/api/user"
    import { StatusNotOkError } from "$lib/api/request"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        try {
            await login(userForm.getUsername(), userForm.getPassword())
            dispatch("userLoggedIn")
        } catch (err) {
            if (err instanceof StatusNotOkError) {
                error = err.message
                return
            }

            // TODO: centralize all errors
            error =
                "An error has occurred while logging in, unable to connect to server"
        }
    }
</script>

<UserForm bind:this={userForm} {error} {onClick} buttonLabel={"Login"} />
