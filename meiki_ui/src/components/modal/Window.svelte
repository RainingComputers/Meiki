<script lang="ts">
    export let width: string
</script>

<div
    class="w-96 max-h-fit m-auto bg-base-1 border-base-2 rounded-xl flex flex-col py-6 px-6"
    style="width: {width};"
>
    <slot />
</div>
