<script lang="ts">
    import Error from "$cmp/toast/Error.svelte"
    export let error: string = ""
</script>

<div
    class="w-96 max-h-fit m-auto bg-panel border-borderPanel rounded-xl flex flex-col py-6 px-6"
>
    <div class="flex flex-col gap-5">
        {#if error}
            <Error fullWidth={true}>{error}</Error>
        {/if}
        <slot />
    </div>
</div>
