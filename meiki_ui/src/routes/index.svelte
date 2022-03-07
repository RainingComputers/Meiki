<!--  TODO: index page, check if the user is already looged,
    if then show hero page else
     login buttton -->
<script lang="ts">
    import { goto } from "$app/navigation"
    import { onMount } from "svelte"

    import { authStatus } from "../api/authStatus"

    onMount(async () => {
        try {
            await authStatus()
        } catch {
            goto(`/login`)
        }
    })
</script>

<h1>Meiki UI</h1>
