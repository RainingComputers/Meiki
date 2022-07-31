<script lang="ts">
    import Window from "$cmp/modal/Window.svelte"
    import Error from "$cmp/toast/Error.svelte"
    export let error: string = ""
</script>

<Window width="24rem">
    <div class="flex flex-col gap-5">
        {#if error}
            <Error fullWidth={true}>{error}</Error>
        {/if}
        <slot />
    </div>
</Window>
