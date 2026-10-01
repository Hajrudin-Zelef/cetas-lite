---
id: collect-261001-general-networking/general-networking/development-examples-helloworld-html-da9a1fe4-2
title: "Hello world module & pluginï"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-general-networking/development-examples-helloworld-html-da9a1fe4.md
source_anchor: ""
source_lines: [162, 363]
sha256: 2ecceb665eb1e173a08589bcf4388a250a32deea4e533b1fbe1465a00f2c8e94
---

# Hello world module & pluginï

```
<items>
    <!-- container -->
    <general>
        <!-- fields -->
        <Enabled type="BooleanField">
            <default>1</default>
            <Required>Y</Required>
        </Enabled>
        <SMTPHost type="NetworkField">
            <Required>Y</Required>
        </SMTPHost>
        <FromEmail type="EmailField">
            <default>sample@example.com</default>
            <Required>Y</Required>
        </FromEmail>
        <ToEmail type="EmailField">
            <Required>Y</Required>
        </ToEmail>
        <Description type="TextField">
            <Required>Y</Required>
        </Description>
    </general>
</items>
â¦â¦â¦
```
All available field types can be found in the models/OPNsense/Base/FieldTypes directory. If specific field types support additional parameters, for example for validation, they should be registered in the model as well (just like the default tag in Enabled).

### Presentation XMLï

Create a presentation XML to feed your template

Because creating forms is one of the key assets of the system, we have build some easy to use wrappers to guide you through the process. First we create an XML file for the presentation, which defines fields to use and adds some information for your template to render. Create a file in your controller directory using the sub directory forms and name it general.xml. Next copy in the following content:

```
<form>
    <field>
        <id>helloworld.general.Enabled</id>
        <label>enabled</label>
        <type>checkbox</type>
        <help>Enable this feature</help>
    </field>
    <field>
        <id>helloworld.general.SMTPHost</id>
        <label>SMTPHost</label>
        <type>text</type>
        <help><![CDATA[ip address of the mail host]]></help>
        <hint>choose a valid IPv4/v6 address</hint>
    </field>
    <field>
        <id>helloworld.general.FromEmail</id>
        <label>Email (from)</label>
        <type>text</type>
    </field>
    <field>
        <id>helloworld.general.ToEmail</id>
        <label>Email (to)</label>
        <type>text</type>
    </field>
    <field>
        <id>helloworld.general.Description</id>
        <label>Description</label>
        <type>text</type>
    </field>
 </form>
```
All items should contain at least an id (where to map data from/to), a type (how to display) and a label, which identifies it to the user. Optional you may add additional fields like help or mark features as being only for advanced users. (The Volt template defines which attributes are usable.)

Now we need to tell the controller to use this information and pass it to your template, so change the IndexController.php and add this line:

```
$this->view->generalForm = $this->getForm("general");
```
And we are ready to update the (Volt) template with this information. Letâs remove the â<h1>Hello World!</h1>â line and replace it with something like this:

```
{{ partial("layout_partials/base_form",['fields':generalForm,'id':'frm_GeneralSettings'])}}
```
This tells the template system to add a form using the contents of generalForm and name it frm_GeneralSettings in the HTML page. Based on a standard template part which is already part of the standard system, named base_form.volt.

When opening the page again it will render like this:

### Create API callsï

Create API calls to retrieve and store data

The framework provides some helpful utilities to get and set data from
and to the configuration XML by using your defined model. First step in
binding your model to the system is to point the `SettingsController` to the model and teach it how
it should return the data.  For this we add two lines to the controller created earlier:

```
class SettingsController extends ApiMutableModelControllerBase
 {
     protected static $internalModelClass = 'OPNsense\HelloWorld\HelloWorld';
     protected static $internalModelName = 'helloworld';
     public function getAction()
     {
         $data = parent::getAction();
         $data[self::$internalModelName]['general']['%ToEmail'] = gettext('Enter recipient here');
         return $data;
     }
 }
```
Note

The `getAction()` function can add dynamic extra information to data fetches. When we get to fetching the data
this will be further explained.

The `$internalModelClass` creates the model for you, so you donât have to create one manually (and define get and
set actions), `$internalModelName` names the response container.

Similarly, we will do this for the `ServiceController` as well

```
class ServiceController extends ApiMutableServiceControllerBase
 {
     protected static $internalServiceClass = 'OPNsense\HelloWorld\HelloWorld';
     protected static $internalServiceClass = 'helloworld';
 }
```
You can test the result (while logged in as root), by going to this address:

```
http[s]://<your ip>/api/helloworld/settings/get
```
Which will output a json structure in your browser like:

```
{
    "helloworld": {
        "general": {
        "Enabled": "1",
        "SMTPHost": "",
        "FromEmail": "sample@example.com",
        "ToEmail": "",
        "%ToEmail": "Enter recipient here",
        "Description": ""
        }
    }
}
```
Note

The outer container is named âhelloworldâ and contains all fields defined in the model.

Note

`ApiMutableModelControllerBase` contains more shared functionality for grid like operations as well, most of
our api controllers use this as a base. When interested in the internals of the get action itself, search
for `getAction` in `ApiMutableModelControllerBase.php`.

Note

`%ToEmail` is similar to `<hint>Enter recipient here</hint>` except that it can be used to generate
a dynamic hint during data fetch. Fields with a `%` prefix cannot be stored and are not relevant to the
model data, but do refer to the respective field type without the prefix.

### Support jQuery API callsï

Update the view to support the API calls using jQuery

Now we need to link the events to the backend code to be able to load and save our form, by using the OPNsense libraries you can validate your data automatically.

Add this to the index.volt template from the HelloWorld module:

```
<script type="text/javascript">
    $( document ).ready(function() {
        mapDataToFormUI({'frm_GeneralSettings':"/api/helloworld/settings/get"}).done(function(data){
            // place actions to run after load, for example update form styles.
        });
        // link save button to API set action
        $("#saveAct").click(function(){
            saveFormToEndpoint("/api/helloworld/settings/set",'frm_GeneralSettings',function(){
                // action to run after successful save, for example reconfigure service.
            });
        });
    });
</script>
<div class="col-md-12">
    <button class="btn btn-primary"  id="saveAct" type="button"><b>{{ lang._('Save') }}</b></button>
</div>
```
The first piece of javascript code handles the loading of data when opening the form, then a button is linked to the save event.

Letâs give it a try and save our data, without modifying it first.

Next correct the errors and save again, on successful save the data should be stored in the config.xml. If you want to change validation messages, just edit the model XML and add your message in the ValidationMessage tag. For example:

```
<ToEmail type="EmailField">
    <Required>Y</Required>
    <ValidationMessage>Please specify a valid email address.</ValidationMessage>
</ToEmail>
```
Changes the âEmail address is invalid.â into âPlease specify a valid email address.â when an invalid email address is provided. Since this field is required the validation message for an empty field will always be âA value is required.â instead.

### Add actionsï

Add some activity to the module

