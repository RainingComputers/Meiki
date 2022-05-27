<script lang="ts">
    import { createEventDispatcher } from "svelte"
    import { login } from "$lib/api/user"
    import { formatRequestError } from "$lib/api/request"
    import UserForm from "./UserForm.svelte"

    let error: string = ""
    const dispatch = createEventDispatcher()

    let userForm: UserForm

    async function onClick() {
        try {
            await login(userForm.getUsername(), userForm.getPassword())
            dispatch("userLoggedIn")
        } catch (err) {
            error = formatRequestError(err, "logging in")
        }
    }
</script>

<UserForm bind:this={userForm} {error} {onClick} buttonLabel={"Login"} />
