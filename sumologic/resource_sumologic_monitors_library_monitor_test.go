package sumologic

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccSumologicMonitorsLibraryMonitor_basic(t *testing.T) {
	var monitorsLibraryMonitor MonitorsLibraryMonitor
	testModifiedAt := "2020-08-03T16:46:19Z"
	testCreatedBy := "createdBy-f2f1f"
	testIsLocked := false
	testMonitorType := "monitorType-xmUaG"
	testIsSystem := false
	testName := "name-JMH7N"
	testParentId := "parentId-irI70"
	testIsMutable := false
	testNotifications := []string{"\"notifications-TBqYJ\""}
	testType := "type-z0sN2"
	testVersion := 0
	testTriggers := []string{"\"triggers-lfpXf\""}
	testQueries := []string{"\"queries-e8uRI\""}
	testDescription := "description-JHJLo"
	testModifiedBy := "modifiedBy-3IdrZ"
	testCreatedAt := "2020-08-03T16:46:19Z"
	testContentType := "contentType-3Br1e"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorsLibraryMonitorDestroy(monitorsLibraryMonitor),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSumologicMonitorsLibraryMonitorConfigImported(testModifiedAt, testCreatedBy, testIsLocked, testMonitorType, testIsSystem, testName, testParentId, testIsMutable, testNotifications, testType, testVersion, testTriggers, testQueries, testDescription, testModifiedBy, testCreatedAt, testContentType),
			},
			{
				ResourceName:      "sumologic_monitor.foo",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
func TestAccMonitorsLibraryMonitor_create(t *testing.T) {
	var monitorsLibraryMonitor MonitorsLibraryMonitor
	testModifiedAt := "2020-08-03T16:46:19Z"
	testCreatedBy := "createdBy-c0Ozz"
	testIsLocked := false
	testMonitorType := "monitorType-5ZjSB"
	testIsSystem := false
	testName := "name-lOYzt"
	testParentId := "parentId-SAzzN"
	testIsMutable := false
	testNotifications := []string{"\"notifications-ISRrx\""}
	testType := "type-DavFz"
	testVersion := 0
	testTriggers := []string{"\"triggers-IvzGU\""}
	testQueries := []string{"\"queries-38tkL\""}
	testDescription := "description-BtcFH"
	testModifiedBy := "modifiedBy-MMHG8"
	testCreatedAt := "2020-08-03T16:46:19Z"
	testContentType := "contentType-cTsuR"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorsLibraryMonitorDestroy(monitorsLibraryMonitor),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicMonitorsLibraryMonitor(testModifiedAt, testCreatedBy, testIsLocked, testMonitorType, testIsSystem, testName, testParentId, testIsMutable, testNotifications, testType, testVersion, testTriggers, testQueries, testDescription, testModifiedBy, testCreatedAt, testContentType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorsLibraryMonitorExists("sumologic_monitor.test", &monitorsLibraryMonitor, t),
					testAccCheckMonitorsLibraryMonitorAttributes("sumologic_monitor.test"),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_at", testModifiedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_by", testCreatedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_locked", strconv.FormatBool(testIsLocked)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "monitor_type", testMonitorType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_system", strconv.FormatBool(testIsSystem)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "name", testName),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "parent_id", testParentId),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_mutable", strconv.FormatBool(testIsMutable)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "notifications.0", strings.Replace(testNotifications[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "type", testType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "version", strconv.Itoa(testVersion)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "triggers.0", strings.Replace(testTriggers[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "queries.0", strings.Replace(testQueries[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "description", testDescription),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_by", testModifiedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_at", testCreatedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "content_type", testContentType),
				),
			},
		},
	})
}

func TestAccMonitorsLibraryMonitor_update(t *testing.T) {
	var monitorsLibraryMonitor MonitorsLibraryMonitor
	testModifiedAt := "2020-08-03T16:46:19Z"
	testCreatedBy := "createdBy-Du0BY"
	testIsLocked := false
	testMonitorType := "monitorType-93CaT"
	testIsSystem := false
	testName := "name-NFcxc"
	testParentId := "parentId-fszzK"
	testIsMutable := false
	testNotifications := []string{"\"notifications-SgrXE\""}
	testType := "type-4AD9v"
	testVersion := 0
	testTriggers := []string{"\"triggers-CsPO9\""}
	testQueries := []string{"\"queries-msvx8\""}
	testDescription := "description-aTsWM"
	testModifiedBy := "modifiedBy-IiNRv"
	testCreatedAt := "2020-08-03T16:46:19Z"
	testContentType := "contentType-ju7xn"

	testUpdatedModifiedAt := "2020-08-03T16:46:19ZUpdate"
	testUpdatedCreatedBy := "createdBy-ue0hwUpdate"
	testUpdatedIsLocked := false
	testUpdatedMonitorType := "monitorType-pLijAUpdate"
	testUpdatedIsSystem := false
	testUpdatedName := "name-z7WzNUpdate"
	testUpdatedParentId := "parentId-bSTPCUpdate"
	testUpdatedIsMutable := false
	testUpdatedNotifications := []string{"\"notifications-U3cj5\""}
	testUpdatedType := "type-uz9m2Update"
	testUpdatedVersion := 1
	testUpdatedTriggers := []string{"\"triggers-kQuSS\""}
	testUpdatedQueries := []string{"\"queries-GP6Id\""}
	testUpdatedDescription := "description-YZ1GsUpdate"
	testUpdatedModifiedBy := "modifiedBy-g8kbKUpdate"
	testUpdatedCreatedAt := "2020-08-03T16:46:19ZUpdate"
	testUpdatedContentType := "contentType-psWgKUpdate"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorsLibraryMonitorDestroy(monitorsLibraryMonitor),
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicMonitorsLibraryMonitor(testModifiedAt, testCreatedBy, testIsLocked, testMonitorType, testIsSystem, testName, testParentId, testIsMutable, testNotifications, testType, testVersion, testTriggers, testQueries, testDescription, testModifiedBy, testCreatedAt, testContentType),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckMonitorsLibraryMonitorExists("sumologic_monitor.test", &monitorsLibraryMonitor, t),
					testAccCheckMonitorsLibraryMonitorAttributes("sumologic_monitor.test"),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_at", testModifiedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_by", testCreatedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_locked", strconv.FormatBool(testIsLocked)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "monitor_type", testMonitorType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_system", strconv.FormatBool(testIsSystem)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "name", testName),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "parent_id", testParentId),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_mutable", strconv.FormatBool(testIsMutable)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "notifications.0", strings.Replace(testNotifications[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "type", testType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "version", strconv.Itoa(testVersion)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "triggers.0", strings.Replace(testTriggers[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "queries.0", strings.Replace(testQueries[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "description", testDescription),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_by", testModifiedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_at", testCreatedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "content_type", testContentType),
				),
			},
			{
				Config: testAccSumologicMonitorsLibraryMonitorUpdate(testUpdatedModifiedAt, testUpdatedCreatedBy, testUpdatedIsLocked, testUpdatedMonitorType, testUpdatedIsSystem, testUpdatedName, testUpdatedParentId, testUpdatedIsMutable, testUpdatedNotifications, testUpdatedType, testUpdatedVersion, testUpdatedTriggers, testUpdatedQueries, testUpdatedDescription, testUpdatedModifiedBy, testUpdatedCreatedAt, testUpdatedContentType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_at", testUpdatedModifiedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_by", testUpdatedCreatedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_locked", strconv.FormatBool(testUpdatedIsLocked)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "monitor_type", testUpdatedMonitorType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_system", strconv.FormatBool(testUpdatedIsSystem)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "name", testUpdatedName),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "parent_id", testUpdatedParentId),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "is_mutable", strconv.FormatBool(testUpdatedIsMutable)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "notifications.0", strings.Replace(testUpdatedNotifications[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "type", testUpdatedType),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "version", strconv.Itoa(testUpdatedVersion)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "triggers.0", strings.Replace(testUpdatedTriggers[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "queries.0", strings.Replace(testUpdatedQueries[0], "\"", "", 2)),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "description", testUpdatedDescription),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "modified_by", testUpdatedModifiedBy),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "created_at", testUpdatedCreatedAt),
					resource.TestCheckResourceAttr("sumologic_monitor.test", "content_type", testUpdatedContentType),
				),
			},
		},
	})
}

func testAccCheckMonitorsLibraryMonitorDestroy(monitorsLibraryMonitor MonitorsLibraryMonitor) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*Client)
		for _, r := range s.RootModule().Resources {
			id := r.Primary.ID
			u, err := client.MonitorsRead(id)
			if err != nil {
				return fmt.Errorf("Encountered an error: " + err.Error())
			}
			if u != nil {
				return fmt.Errorf("MonitorsLibraryMonitor %s still exists", id)
			}
		}
		return nil
	}
}

func testAccCheckMonitorsLibraryMonitorExists(name string, monitorsLibraryMonitor *MonitorsLibraryMonitor, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			//need this so that we don't get an unused import error for strconv in some cases
			return fmt.Errorf("Error = %s. MonitorsLibraryMonitor not found: %s", strconv.FormatBool(ok), name)
		}

		//need this so that we don't get an unused import error for strings in some cases
		if strings.EqualFold(rs.Primary.ID, "") {
			return fmt.Errorf("MonitorsLibraryMonitor ID is not set")
		}

		id := rs.Primary.ID
		client := testAccProvider.Meta().(*Client)
		newMonitorsLibraryMonitor, err := client.MonitorsRead(id)
		if err != nil {
			return fmt.Errorf("MonitorsLibraryMonitor %s not found", id)
		}
		monitorsLibraryMonitor = newMonitorsLibraryMonitor
		return nil
	}
}
func testAccCheckSumologicMonitorsLibraryMonitorConfigImported(modifiedAt string, createdBy string, isLocked bool, monitorType string, isSystem bool, name string, parentId string, isMutable bool, notifications []string, type_field string, version int, triggers []string, queries []string, description string, modifiedBy string, createdAt string, contentType string) string {
	return fmt.Sprintf(`
resource "sumologic_monitor" "foo" {
      modified_at = "%s"
      created_by = "%s"
      is_locked = %t
      monitor_type = "%s"
      is_system = %t
      name = "%s"
      parent_id = "%s"
      is_mutable = %t
      notifications = %v
      type = "%s"
      version = %d
      triggers = %v
      queries = %v
      description = "%s"
      modified_by = "%s"
      created_at = "%s"
      content_type = "%s"
}
`, modifiedAt, createdBy, isLocked, monitorType, isSystem, name, parentId, isMutable, notifications, type_field, version, triggers, queries, description, modifiedBy, createdAt, contentType)
}

func testAccSumologicMonitorsLibraryMonitor(modifiedAt string, createdBy string, isLocked bool, monitorType string, isSystem bool, name string, parentId string, isMutable bool, notifications []string, type_field string, version int, triggers []string, queries []string, description string, modifiedBy string, createdAt string, contentType string) string {
	return fmt.Sprintf(`
resource "sumologic_monitor" "test" {
    modified_at = "%s"
    created_by = "%s"
    is_locked = %t
    monitor_type = "%s"
    is_system = %t
    name = "%s"
    parent_id = "%s"
    is_mutable = %t
    notifications = %v
    type = "%s"
    version = %d
    triggers = %v
    queries = %v
    description = "%s"
    modified_by = "%s"
    created_at = "%s"
    content_type = "%s"
}
`, modifiedAt, createdBy, isLocked, monitorType, isSystem, name, parentId, isMutable, notifications, type_field, version, triggers, queries, description, modifiedBy, createdAt, contentType)
}

func testAccSumologicMonitorsLibraryMonitorUpdate(modifiedAt string, createdBy string, isLocked bool, monitorType string, isSystem bool, name string, parentId string, isMutable bool, notifications []string, type_field string, version int, triggers []string, queries []string, description string, modifiedBy string, createdAt string, contentType string) string {
	return fmt.Sprintf(`
resource "sumologic_monitor" "test" {
      modified_at = "%s"
      created_by = "%s"
      is_locked = %t
      monitor_type = "%s"
      is_system = %t
      name = "%s"
      parent_id = "%s"
      is_mutable = %t
      notifications = %v
      type = "%s"
      version = %d
      triggers = %v
      queries = %v
      description = "%s"
      modified_by = "%s"
      created_at = "%s"
      content_type = "%s"
}
`, modifiedAt, createdBy, isLocked, monitorType, isSystem, name, parentId, isMutable, notifications, type_field, version, triggers, queries, description, modifiedBy, createdAt, contentType)
}

func testAccCheckMonitorsLibraryMonitorAttributes(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		f := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet(name, "modified_at"),
			resource.TestCheckResourceAttrSet(name, "created_by"),
			resource.TestCheckResourceAttrSet(name, "is_locked"),
			resource.TestCheckResourceAttrSet(name, "monitor_type"),
			resource.TestCheckResourceAttrSet(name, "is_system"),
			resource.TestCheckResourceAttrSet(name, "name"),
			resource.TestCheckResourceAttrSet(name, "parent_id"),
			resource.TestCheckResourceAttrSet(name, "is_mutable"),
			resource.TestCheckResourceAttrSet(name, "notifications"),
			resource.TestCheckResourceAttrSet(name, "type"),
			resource.TestCheckResourceAttrSet(name, "version"),
			resource.TestCheckResourceAttrSet(name, "triggers"),
			resource.TestCheckResourceAttrSet(name, "queries"),
			resource.TestCheckResourceAttrSet(name, "description"),
			resource.TestCheckResourceAttrSet(name, "modified_by"),
			resource.TestCheckResourceAttrSet(name, "created_at"),
			resource.TestCheckResourceAttrSet(name, "content_type"),
		)
		return f(s)
	}
}
