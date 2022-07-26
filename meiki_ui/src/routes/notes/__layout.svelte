<script lang="ts" context="module">
    import { authStatus } from "$lib/api/user"
    import type { Load } from "@sveltejs/kit"

    export const load: Load = async () => {
        try {
            await authStatus()
            return {}
        } catch {
            return {
                status: 302,
                redirect: "/login",
            }
        }
    }
</script>

<slot />
