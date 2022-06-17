<script lang="ts">
    import Error from "$cmp/toast/Error.svelte"

    export let message: string
</script>

<span class="flex px-4 py-4 w-full">
    <Error fullWidth>{message}</Error>
</span>
