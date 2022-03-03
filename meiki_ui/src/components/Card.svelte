<script lang="ts">
    export let width: number = 400
</script>

<span
    class="flex flex-col bg-gray-50 border-gray-200 border p-5 drop-shadow-card gap-10 items-center"
    style="width: {width}px;"
>
    <slot />
</span>
