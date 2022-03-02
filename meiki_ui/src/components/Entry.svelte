<script lang="ts">
    export let label: string = "label"
</script>

<div class=" flex flex-col w-full gap-1">
    <span class=" text-sm">{label}</span>
    <input class=" bg-gray-300 rounded-lg p-2" type="text" required />
</div>
