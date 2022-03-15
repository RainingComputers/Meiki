<script lang="ts">
    // TODO: Add input validation
    // TODO: show user exists
    // TODO: show better error messages

    import { createEventDispatcher } from "svelte"
    import { login } from "$lib/api/login"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        try {
            await login(userForm.getUsername(), userForm.getPassword())

            dispatch("userLoggedIn")
        } catch (err) {
            error = "An error has occured while logging in, try again"
        }
    }
</script>

<UserForm bind:this={userForm} {error} {onClick} buttonLabel={"Login"} />
