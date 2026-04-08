# Minutes of Meeting (MoM)

**Project Title:** Blueframe SDK Development – COMTRADE Parser Exploration  
**Team Number:** 03-SEL-BFSDK  
**Client / Sponsor:** Schweitzer Engineering Laboratories (SEL)  
**Mentor(s):** Joe Stanley  
**Date:** April 7  
**Time:** 2:00 PM  
**Location / Platform:** In-person, Spark  
**Participants (Team & Client):** Genevieve Kochel, Darron Li, Lucas Phillips, Caleb Hansen, Joe Stanley, Jedd Bartlet  
**Meeting Number / Version:** 2

## 1. Agenda
- **Item 1:** Introduction to COMTRADE File Parsing Needs  
- **Item 2:** Overview of Proposed GO Program  
- **Item 3:** Align on Project Scope and Deliverables  
- **Item 4:** Discuss Integration with Existing Blueframe SDK

## 2. Key Discussion Points
- **COMTRADE Background:**  
  Jedd Bartlet highlighted the need for a reliable program to parse COMTRADE files, which contain substation disturbance and event data. Accurate parsing is critical for analyzing events, validating device performance, and supporting WAMS analytics.

- **Proposed GO Program:**  
  - Written in GO to leverage performance and cross-platform capabilities.  
  - Provides functions to read, validate, and extract signal data from COMTRADE files (.cfg and .dat).  
  - Supports conversion to data structures usable in Python SDK and other Blueframe applications.  
  - Initial scope includes basic file reading, error detection, and sample extraction; advanced analytics and aggregation planned for future iterations.

- **Integration Considerations:**  
  - Program output must be compatible with the Python SDK abstraction layer.  
  - Consider creating a REST wrapper around GO parser to allow SDK calls directly.  
  - Must maintain reliability and error handling standards consistent with Blueframe high-availability environment.

- **Value Proposition:**  
  - Engineers can analyze disturbance data programmatically without manual COMTRADE file inspection.  
  - Enables automation for device validation and fault detection pipelines within Blueframe.

## 3. Decisions Made
- Proceed with developing a GO-based COMTRADE parser.  
- Program will initially support reading and extracting key signal and timestamp data.  
- Integration pathway with Python SDK confirmed via a REST interface approach.

## 4. Action Items / Responsibilities
- **Task:** Define required COMTRADE fields and validation rules  
  **Assigned To:** Jedd Bartlet, Darron Li  
  **Deadline:** April 21  
  **Priority:** High

- **Task:** Prototype GO parser for single-file extraction  
  **Assigned To:** Jedd Bartlet, Lucas Phillips  
  **Deadline:** April 28  
  **Priority:** High

- **Task:** Draft integration approach with Python SDK  
  **Assigned To:** Genevieve Kochel, Caleb Hansen  
  **Deadline:** May 5  
  **Priority:** Medium

- **Task:** Document parser API and usage guidelines  
  **Assigned To:** Entire Team  
  **Deadline:** May 12  
  **Priority:** Medium

## 5. Client Feedback / Clarifications
- Parsing must handle large COMTRADE files efficiently and report errors clearly.  
- Outputs should match engineering workflows and integrate seamlessly with existing Blueframe tools.  
- Reliability and cross-platform compatibility emphasized.

## 6. Linkage to Deliverables (optional)
- **Relevant Requirement Document Section(s):** Data Analysis Tools; COMTRADE Integration  
- **Impact on Sprint / Milestone:** Adds new toolchain requirement for Sprint 2  
- **Presentation / Report Updates Needed:** Include GO parser design and COMTRADE workflow diagram

## 7. Next Steps & Follow-Up
- **Deliverables before next meeting:** Field requirements list, GO parser prototype, preliminary integration outline  
- **Next meeting scheduled on:** April 28  
- **Agreed communication channel:** Teams / Email

**Prepared By:** Caleb Hansen  
**Date of Circulation:** April 7
