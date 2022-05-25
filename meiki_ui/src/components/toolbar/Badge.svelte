<script lang="ts">
    export let label: string = ""
</script>

<span
    class="flex text-gray-200 font-light text-xs w-24 overflow-hidden px-3 gap-1 items-center"
>
    <span class="stroke-gray-200 h-3 w-3">
        <slot />
    </span>
    {label}
</span>
