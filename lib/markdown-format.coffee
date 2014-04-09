Spawn = require('child_process').spawn

class MarkdownFormat
  constructor: ->
    atom.workspace.eachEditor (editor) =>
      @handleSave(editor)

  destroy: ->

  handleSave: (editor) ->
    buffer = editor.getBuffer()
    buffer.on 'saved', =>
      if editor.getGrammar().scopeName is 'source.gfm'
        options = [
          '-w'
          editor.getUri()
        ]
        markdownfmt = Spawn(atom.config.get('markdown-format.pathToBinary'), options)

module.exports =
  configDefaults:
    pathToBinary: 'markdownfmt'

  activate: ->
    @markdownFormat = new MarkdownFormat()

  deactivate: ->
    @markdownFormat.destroy()
