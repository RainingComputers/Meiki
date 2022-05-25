<script lang="ts">
    export let title: string
</script>

<span
    class="flex flex-row flex-grow text-gray-200 justify-center font-light items-center"
>
    {title}
    <slot />
</span>
