<script>
    export let fullWidth = false
</script>

<span
    class=" text-red-500 font-bold bg-red-50 p-2 rounded-lg border-2 border-red-400 drop-shadow-card"
    class:w-full={fullWidth}
>
    <slot />
</span>
