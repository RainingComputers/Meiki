<script lang="ts">
    import Markdoc from "@markdoc/markdoc"

    export let text: string = ""

    $: ast = Markdoc.parse(text)
    $: content = Markdoc.transform(ast)
</script>

<p data-cy="renderer" class="py-3 px-5 pb-48 overflow-scroll h-full prose max-w-full">
    {@html Markdoc.renderers.html(content)}
</p>
