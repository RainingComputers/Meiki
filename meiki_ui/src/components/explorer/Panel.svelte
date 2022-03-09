<script lang="ts">
    export let width: string = "24rem"
</script>

<div
    class=" bg-gray-50 border-gray-200 border flex flex-col gap-1 py-2"
    style="width: {width}"
>
    <slot />
</div>
